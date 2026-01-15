// const apiUrl = process.env.NEXT_PUBLIC_API_URL ?? "";
const apiUrl = "https://click-api.y3eet.me";

if (!apiUrl) {
  console.warn("Environment variable NEXT_PUBLIC_API_URL is not set.");
}

export const config = { apiUrl };
