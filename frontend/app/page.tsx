import { createOrder } from "@/lib/orders";

export default function Home() {
  async function testCreate() {
    try {
      const res = await createOrder("order-ui-1", "user-1");
      console.log(res);
      alert("Order created!");
    } catch (e: any) {
      alert(e.message);
    }
  }

  return (
    <main className="p-8">
      <h1 className="text-3xl font-bold">FoodOS</h1>
      <button
        onClick={testCreate}
        className="mt-6 px-4 py-2 bg-black text-white rounded"
      >
        Create Test Order
      </button>
    </main>
  );
}
