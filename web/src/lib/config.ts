const apiUrl = process.env.NEXT_PUBLIC_API_URL ?? "";

if (!apiUrl) {
  console.warn("Environment variable NEXT_PUBLIC_API_URL is not set.");
}

export const config = { apiUrl };
