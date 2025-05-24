import { useEffect, useState } from "react";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "./ui/table";
import { differenceInDays, parseISO } from "date-fns";

interface User {
  name: string;
  createDate: string;
  passwordChangedDate: string;
  lastAccessDate: string;
  mfaEnabled: boolean;
}

export function UserTable() {
  const [users, setUsers] = useState<User[]>([]);
  const [filter, setFilter] = useState<"all" | "true" | "false">("all");
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState<string | null>(null);

  useEffect(() => {
    fetch(`http://localhost:8080/api/users`)
      .then((res) => {
        if (!res.ok) throw new Error(`Server responded ${res.status}`);
        return res.json() as Promise<User[]>;
      })
      .then(setUsers)
      .catch((err) => setError(err.message))
      .finally(() => setLoading(false));
  }, []);

  if (loading) return <p>Loading users…</p>;
  if (error) return <p className="text-red-600">Error: {error}</p>;

  const today = new Date();
  const filtered = users.filter((u) => {
    if (filter === "all") return true;
    return filter === "true" ? u.mfaEnabled : !u.mfaEnabled;
  });

  return (
    <>
      <div className="mb-4 flex items-center space-x-2">
        <label htmlFor="mfa" className="font-medium">
          Filter by MFA:
        </label>
        <select
          id="mfa"
          className="border rounded px-2 py-1"
          value={filter}
          onChange={(e) =>
            setFilter(e.target.value as "all" | "true" | "false")
          }
        >
          <option value="all">All</option>
          <option value="true">Enabled</option>
          <option value="false">Disabled</option>
        </select>
      </div>

      <Table>
        <TableCaption>Users with live “days since” metrics</TableCaption>
        <TableHeader>
          <TableRow>
            <TableHead>Name</TableHead>
            <TableHead>Create Date</TableHead>
            <TableHead>Password Changed</TableHead>
            <TableHead>Days Since Pwd Change</TableHead>
            <TableHead>Last Access</TableHead>
            <TableHead>Days Since Access</TableHead>
            <TableHead>MFA Enabled</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {filtered.map((u) => {
            const pwdDays = differenceInDays(
              today,
              parseISO(u.passwordChangedDate)
            );
            const accessDays = differenceInDays(
              today,
              parseISO(u.lastAccessDate)
            );
            const rowClass = accessDays > 90 ? "bg-red-100" : "";

            return (
              <TableRow key={u.name} className={rowClass}>
                <TableCell className="font-medium">{u.name}</TableCell>
                <TableCell>{u.createDate}</TableCell>
                <TableCell>{u.passwordChangedDate}</TableCell>
                <TableCell>{pwdDays}</TableCell>
                <TableCell>{u.lastAccessDate}</TableCell>
                <TableCell>{accessDays}</TableCell>
                <TableCell>{u.mfaEnabled ? "Yes" : "No"}</TableCell>
              </TableRow>
            );
          })}
        </TableBody>
        <TableFooter>
          <TableRow>
            <TableCell colSpan={6}>Total users shown</TableCell>
            <TableCell className="text-right">{filtered.length}</TableCell>
          </TableRow>
        </TableFooter>
      </Table>
    </>
  );
}
