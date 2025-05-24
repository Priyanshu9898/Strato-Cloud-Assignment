import React from "react";
import { render, screen, waitFor, fireEvent } from "@testing-library/react";
import { beforeEach, afterEach, describe, it, expect, vi } from "vitest";
import { UserTable } from "./UserTable";

// --- mock data ------------------------------------------------------
const mockUsers = [
  {
    name: "UserA",
    createDate: "2021-01-01",
    passwordChangedDate: "2024-05-01",
    lastAccessDate: "2025-01-01",
    mfaEnabled: true,
  },
  {
    name: "UserB",
    createDate: "2022-01-01",
    passwordChangedDate: "2023-01-01",
    lastAccessDate: "2024-01-01",
    mfaEnabled: false,
  },
];

describe("UserTable", () => {
  beforeEach(() => {
    vi.spyOn(global, "fetch").mockResolvedValueOnce({
      ok: true,
      json: async () => mockUsers,
    } as never);
  });

  afterEach(() => {
    vi.restoreAllMocks();
  });

  it("shows loading, then renders all users with correct MFA text", async () => {
    render(<UserTable />);

    // Loading state
    expect(screen.getByText(/loading users/i)).toBeInTheDocument();

    // Wait for data
    await waitFor(() => screen.getByText("UserA"));

    expect(screen.getByText("UserA")).toBeInTheDocument();
    expect(screen.getByText("UserB")).toBeInTheDocument();
    expect(screen.getByText("Yes")).toBeInTheDocument();
    expect(screen.getByText("No")).toBeInTheDocument();
  });

  it("filters by MFA status correctly", async () => {
    render(<UserTable />);
    await waitFor(() => screen.getByText("UserA"));

    // Both present
    expect(screen.getByText("UserA")).toBeInTheDocument();
    expect(screen.getByText("UserB")).toBeInTheDocument();

    // Filter to only MFA-disabled
    fireEvent.change(screen.getByLabelText(/filter by mfa/i), {
      target: { value: "false" },
    });
    expect(screen.queryByText("UserA")).toBeNull();
    expect(screen.getByText("UserB")).toBeInTheDocument();

    // Filter to only MFA-enabled
    fireEvent.change(screen.getByLabelText(/filter by mfa/i), {
      target: { value: "true" },
    });
    expect(screen.getByText("UserA")).toBeInTheDocument();
    expect(screen.queryByText("UserB")).toBeNull();
  });
});
