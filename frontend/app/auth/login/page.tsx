"use client";

import { useState } from "react";

export default function LoginPage() {
  const [email, setEmail] = useState("");

  async function login() {
    console.log("🔵 Login clicked");
    console.log("🔵 Email:", email);

    try {
      console.log("🟡 Sending fetch request...");

      const res = await fetch("http://127.0.0.1:8081/auth/login", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ email }),
      });

      console.log("🟢 Fetch returned");
      console.log("🟢 Response status:", res.status);

      const text = await res.text();
      console.log("🟢 Raw response text:", text);

      let data;
      try {
        data = JSON.parse(text);
        console.log("🟢 Parsed JSON:", data);
      } catch (e) {
        console.error("🔴 JSON parse failed", e);
        return;
      }

      console.log("🟣 Writing to localStorage...");
      localStorage.setItem("token", data.token);
      localStorage.setItem("role", data.role);

      console.log("🟣 localStorage token:", localStorage.getItem("token"));
      console.log("🟣 localStorage role:", localStorage.getItem("role"));

      console.log("🔴 ABOUT TO REDIRECT");

      if (data.role === "ADMIN") {
        console.log("➡️ Redirecting to /admin/dashboard");
        window.location.assign("/admin/dashboard");
      } else {
        console.log("➡️ Redirecting to /user/dashboard");
        window.location.assign("/user/dashboard");
      }

      console.log("❌ THIS SHOULD NEVER PRINT");

    } catch (err) {
      console.error("🔥 Login exception:", err);
    }
  }

  return (
    <div className="min-h-screen flex items-center justify-center bg-neutral-950">
      <div className="w-96 bg-neutral-900 p-6 rounded-lg">
        <h1 className="text-xl font-bold mb-4">Login</h1>

        <input
          className="w-full p-2 mb-4 bg-neutral-800 rounded"
          placeholder="Email"
          value={email}
          onChange={(e) => setEmail(e.target.value)}
        />

        <button
          onClick={login}
          className="w-full bg-black py-2 rounded"
        >
          Login
        </button>
      </div>
    </div>
  );
}
