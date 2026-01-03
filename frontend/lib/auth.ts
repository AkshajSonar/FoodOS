export function getAuth() {
  if (typeof window === "undefined") return null;

  const token = window.localStorage.getItem("token");
  const role = window.localStorage.getItem("role");

  if (!token || !role) return null;

  return { token, role };
}
