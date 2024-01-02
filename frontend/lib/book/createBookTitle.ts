import { endPoint } from "@/constants";
import axios from "axios";
import { getApiKey } from "../auth/action";

export default async function createBookTitle({
  id,
  name,
  desc,
  categoryIds,
  authorIds,
}: {
  id?: string;
  name: string;
  desc: string;
  categoryIds: string[];
  authorIds: string[];
}) {
  const url = `${endPoint}/v1/booktitles`;

  const data = {
    id: id,
    name: name,
    desc: desc,
    categoryIds: categoryIds,
    authorIds: authorIds,
  };
  console.log(data);
  const token = await getApiKey();
  const headers = {
    accept: "application/json",
    "Content-Type": "application/json",
    Authorization: `Bearer ${token}`,
    // Add other headers as needed
  };

  // Make a POST request with headers
  const res = axios
    .post(url, data, { headers: headers })
    .then((response) => {
      if (response) return response.data;
    })
    .catch((error) => {
      console.error("Error:", error);
      return error.response.data;
    });
  return res;
}