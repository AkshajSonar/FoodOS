type Status =
  | "PLACED"
  | "CONFIRMED"
  | "PREPARING"
  | "PICKED"
  | "DELIVERED"
  | "CANCELLED";

const colors: Record<Status, string> = {
  PLACED: "bg-gray-700 text-white",
  CONFIRMED: "bg-blue-600 text-white",
  PREPARING: "bg-yellow-600 text-black",
  PICKED: "bg-purple-600 text-white",
  DELIVERED: "bg-green-600 text-white",
  CANCELLED: "bg-red-600 text-white",
};

export default function StatusBadge({ status }: { status: Status }) {
  return (
    <span
      className={`px-3 py-1 rounded-full text-sm font-semibold ${colors[status]}`}
    >
      {status}
    </span>
  );
}
