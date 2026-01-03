import Link from "next/link";

export default function UserDashboardPage() {
  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">
        My Dashboard
      </h1>

      <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
        <DashboardCard
          title="My Orders"
          description="View and track your orders"
          href="/user/orders"
        />

        <DashboardCard
          title="Create Order"
          description="Place a new order"
          href="/orders/new"
        />
      </div>
    </div>
  );
}

function DashboardCard({
  title,
  description,
  href,
}: {
  title: string;
  description: string;
  href: string;
}) {
  return (
    <Link
      href={href}
      className="block rounded-lg border border-neutral-800 bg-neutral-900 p-6 hover:bg-neutral-800 transition"
    >
      <h2 className="text-xl font-semibold mb-2">{title}</h2>
      <p className="text-neutral-400">{description}</p>
    </Link>
  );
}
