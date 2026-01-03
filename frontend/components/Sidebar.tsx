import Link from "next/link";

const navItems = [
  { label: "Dashboard", href: "/admin/dashboard" },
  { label: "Orders", href: "/admin/orders" },
  { label: "Users", href: "/admin/users" }, // future
  { label: "Settings", href: "/admin/settings" }, // future
];

export default function Sidebar() {
  return (
    <aside className="w-64 bg-neutral-900 border-r border-neutral-800 p-4">
      <h1 className="text-xl font-bold mb-8">FoodOS</h1>

      <nav className="space-y-2">
        {navItems.map((item) => (
          <Link
            key={item.href}
            href={item.href}
            className="block px-3 py-2 rounded hover:bg-neutral-800 text-sm"
          >
            {item.label}
          </Link>
        ))}
      </nav>
    </aside>
  );
}
