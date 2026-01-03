"use client";

import { updateOrderStatus, getNextStatuses } from "@/lib/orders";
import type { OrderStatus } from "@/lib/orders";
import { useRouter } from "next/navigation";
import { useState } from "react";

export default function AdminControls({
  orderId,
  currentStatus,
}: {
  orderId: string;
  currentStatus: OrderStatus;
}) {
  const router = useRouter();
  const [loading, setLoading] = useState(false);

  const nextStatuses = getNextStatuses(currentStatus);

  async function handleUpdate(status: OrderStatus) {
    try {
      setLoading(true);
      await updateOrderStatus(orderId, status);
      router.refresh(); // re-fetch server data
    } catch {
      alert("Failed to update status");
    } finally {
      setLoading(false);
    }
  }

  if (nextStatuses.length === 0) return null;

  return (
    <div className="mt-10 border-t pt-6">
      <h2 className="text-lg font-semibold mb-3">Admin Controls</h2>

      <div className="flex gap-3 flex-wrap">
        {nextStatuses.map((status) => (
          <button
            key={status}
            disabled={loading}
            onClick={() => handleUpdate(status)}
            className="px-4 py-2 rounded bg-black text-white disabled:opacity-50"
          >
            Move to {status}
          </button>
        ))}
      </div>
    </div>
  );
}
