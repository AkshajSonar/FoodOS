"use client";

import { useEffect, useState } from "react";
import { useRouter } from "next/navigation";
import { getAuth } from "@/lib/auth";

export default function UserLayout({
  children,
}: {
  children: React.ReactNode;
}) {
  const router = useRouter();
  const [ready, setReady] = useState(false);

  useEffect(() => {
    const auth = getAuth();

    if (!auth) {
      router.replace("/auth/login");
      return;
    }

    // 🚫 block admins from user dashboard
    if (auth.role !== "USER") {
      router.replace("/admin/dashboard");
      return;
    }

    setReady(true);
  }, [router]);

  if (!ready) {
    return (
      <div className="min-h-screen flex items-center justify-center">
        Checking user access…
      </div>
    );
  }

  return <>{children}</>;
}
