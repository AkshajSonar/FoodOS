import OrderTimeline, {
  type TimelineStatus,
} from "@/components/OrderTimeline";
import StatusBadge from "@/components/StatusBadge";
import { getOrder } from "@/lib/orders";
import type { OrderStatus } from "@/lib/orders";
import AdminControls from "@/components/AdminControls";

function isTimelineStatus(
  status: OrderStatus
): status is TimelineStatus {
  return status !== "CANCELLED";
}

export default async function OrderTrackingPage({
  params,
}: {
  params: Promise<{ id: string }>;
}) {
  const { id } = await params;
  const order = await getOrder(id);

  let timelineStatus: TimelineStatus | null = null;

  if (isTimelineStatus(order.status)) {
    timelineStatus = order.status;
  }

  return (
    <main className="p-8 max-w-3xl">
      <h1 className="text-3xl font-bold mb-2">
        Order #{order.id}
      </h1>

      <div className="flex items-center gap-3 mt-4">
        <span className="text-gray-400">Current Status:</span>
        <StatusBadge status={order.status} />
      </div>

      {timelineStatus ? (
        <OrderTimeline currentStatus={timelineStatus} />
      ) : (
        <div className="mt-8 p-4 rounded-lg border border-red-600 text-red-500">
          This order was cancelled and will not proceed further.
        </div>
      )}
      <AdminControls
      orderId={order.id}
      currentStatus={order.status}
    />
    </main>
  );
}
