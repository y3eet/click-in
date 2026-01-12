function getEnvVariable(key: string, defaultValue: string = ""): string {
  const value = process.env[key] || defaultValue;
  if (value.length === 0) {
    console.warn(
      `Environment variable ${key} is not set and no default value provided.`
    );
    throw new Error(`Missing environment variable: ${key}`);
  }
  return value;
}

export const config = {
  apiUrl: getEnvVariable("NEXT_PUBLIC_API_URL"),
};
