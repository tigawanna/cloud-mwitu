import { useQuery$ } from "@preact-signals/query";

interface StatusProps {}

export function Status({}: StatusProps) {
  const query = useQuery$(() => ({
    queryKey: ["user"],
    queryFn: fetchUser,
    suspense: true,
    suspenseBehavior: "suspend-eagerly",
  }));
  const data = query.data;
  if (query.error) {
    // const errorMessage = query.error.message;
    return (
      <div
        className="bg-error/10 border border-error text-error px-4 py-3 rounded relative"
        role="alert">
        <strong className="font-bold">Error:</strong>
        {/* @ts-ignore */}
        <span className="block sm:inline">{query.error?.message}</span>
      </div>
    );
  }
  return (
    <div className="w-full h-full flex flex-col items-center justify-center">
      <h1 className={"font-bold text-3xl"}>Status</h1>
      {data.name && <div className={"font-bold text-2xl"}>Name: {data.name}</div>}
    </div>
  );
}

async function fetchUser() {
  const response = await fetch("https://jsonplaceholder.typicode.com/users/1");
  if (!response.ok) {
    throw new Error("Network response was not ok");
  }
  return response.json().catch((error) => {
    throw new Error("JSON parsing error");
  }) as unknown as JsonUser;
}

export interface JsonUser {
  id: number;
  name: string;
  username: string;
  email: string;
  address: Address;
  phone: string;
  website: string;
  company: Company;
}

export interface Address {
  street: string;
  suite: string;
  city: string;
  zipcode: string;
  geo: Geo;
}

export interface Geo {
  lat: string;
  lng: string;
}

export interface Company {
  name: string;
  catchPhrase: string;
  bs: string;
}
