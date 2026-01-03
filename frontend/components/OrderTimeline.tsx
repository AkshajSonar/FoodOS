const STEPS = [
  "PLACED",
  "CONFIRMED",
  "PREPARING",
  "PICKED",
  "DELIVERED",
] as const;

export type TimelineStatus = (typeof STEPS)[number];

export default function OrderTimeline({
  currentStatus,
}: {
  currentStatus: TimelineStatus;
}) {
  const currentIndex = STEPS.indexOf(currentStatus);

  return (
    <div className="flex items-center justify-between mt-8">
      {STEPS.map((step, index) => {
        const isCompleted = index <= currentIndex;

        return (
          <div key={step} className="flex items-center w-full">
            <div
              className={`w-8 h-8 rounded-full flex items-center justify-center text-sm font-bold
              ${isCompleted ? "bg-green-600" : "bg-gray-700"}`}
            >
              {isCompleted ? "✓" : index + 1}
            </div>

            <div className="ml-3 text-sm">
              <p className="font-semibold">{step}</p>
            </div>

            {index < STEPS.length - 1 && (
              <div
                className={`flex-1 h-1 mx-4
                ${isCompleted ? "bg-green-600" : "bg-gray-700"}`}
              />
            )}
          </div>
        );
      })}
    </div>
  );
}
