import OrderCreateForm from "@/components/OrderCreateForm";

export default function NewOrderPage() {
  return (
    <main className="p-8">
      <h1 className="text-3xl font-bold mb-6">New Order</h1>
      <OrderCreateForm />
    </main>
  );
}
