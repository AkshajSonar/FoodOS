import { http } from "./http";

export type OrderStatus =
  | "PLACED"
  | "CONFIRMED"
  | "PREPARING"
  | "PICKED"
  | "DELIVERED"
  | "CANCELLED";

export interface CreateOrderResponse {
  message: string;
  status: OrderStatus;
}

export async function createOrder(
  orderId: string,
  userId: string
): Promise<CreateOrderResponse> {
  return http<CreateOrderResponse>("/orders", {
    method: "POST",
    body: JSON.stringify({
      order_id: orderId,
      user_id: userId,
    }),
  });
}

export async function updateOrderStatus(
  orderId: string,
  status: OrderStatus
) {
  return http(`/orders/${orderId}/status`, {
    method: "PATCH",
    body: JSON.stringify({ status }),
  });
}
export interface Order {
  id: string;
  userId: string;
  status: OrderStatus;
}

export async function getOrder(id: string): Promise<Order> {
  return http<Order>(`/orders/${id}`);
}
export function getNextStatuses(status: OrderStatus): OrderStatus[] {
  switch (status) {
    case "PLACED":
      return ["CONFIRMED", "CANCELLED"];
    case "CONFIRMED":
      return ["PREPARING", "CANCELLED"];
    case "PREPARING":
      return ["PICKED", "CANCELLED"];
    case "PICKED":
      return ["DELIVERED"];
    default:
      return [];
  }
}

