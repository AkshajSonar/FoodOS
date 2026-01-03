import { listOrders } from "@/lib/orders";
import StatusBadge from "@/components/StatusBadge";
import Link from "next/link";

export default async function UserOrdersPage() {
  const orders = await listOrders();

  return (
    <div>
      <h1 className="text-3xl font-bold mb-6">
        My Orders
      </h1>

      <div className="border border-neutral-800 rounded-lg overflow-hidden">
        <table className="min-w-full text-sm">
          <thead className="bg-neutral-900 text-neutral-400">
            <tr>
              <th className="px-4 py-3 text-left">Order ID</th>
              <th className="px-4 py-3 text-left">Status</th>
              <th className="px-4 py-3"></th>
            </tr>
          </thead>
          <tbody>
            {orders.map((o) => (
              <tr key={o.id} className="border-t border-neutral-800">
                <td className="px-4 py-3 font-mono">{o.id}</td>
                <td className="px-4 py-3">
                  <StatusBadge status={o.status} />
                </td>
                <td className="px-4 py-3 text-right">
                  <Link
                    href={`/orders/${o.id}`}
                    className="text-blue-500 hover:underline"
                  >
                    Track
                  </Link>
                </td>
              </tr>
            ))}
          </tbody>
        </table>
      </div>
    </div>
  );
}
