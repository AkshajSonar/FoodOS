"use client";

import { useState } from "react";
import { createOrder } from "@/lib/orders";

export default function OrderCreateForm() {
  const [orderId, setOrderId] = useState("");
  const [loading, setLoading] = useState(false);
  const [message, setMessage] = useState<string | null>(null);
  const [error, setError] = useState<string | null>(null);

  async function handleSubmit(e: React.FormEvent) {
    e.preventDefault();
    setMessage(null);
    setError(null);

    if (!orderId.trim()) {
      setError("Order ID is required");
      return;
    }

    try {
      setLoading(true);
      const res = await createOrder(orderId, "user-1");
      setMessage(`Order created with status: ${res.status}`);
      setOrderId("");
    } catch (err: any) {
      setError(err.message || "Failed to create order");
    } finally {
      setLoading(false);
    }
  }

  return (
    <form
      onSubmit={handleSubmit}
      className="max-w-md rounded-lg border p-6 space-y-4"
    >
      <h2 className="text-xl font-semibold">Create Order</h2>

      <input
        type="text"
        placeholder="Order ID (e.g. order-123)"
        value={orderId}
        onChange={(e) => setOrderId(e.target.value)}
        className="w-full rounded border px-3 py-2"
      />

      <button
        type="submit"
        disabled={loading}
        className="w-full rounded bg-black py-2 text-white disabled:opacity-50"
      >
        {loading ? "Creating..." : "Create Order"}
      </button>

      {message && (
        <p className="text-sm text-green-600">{message}</p>
      )}

      {error && (
        <p className="text-sm text-red-600">{error}</p>
      )}
    </form>
  );
}
