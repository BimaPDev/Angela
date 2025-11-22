import axios from "axios";
import { Platform } from "react-native";

const BASE_URL =
  Platform.OS === "android"
    ? "http://10.0.2.2:8080/api/v1"
    : "http://localhost:8080/api/v1";

const client = axios.create({ baseURL: BASE_URL });

export const api = {
  getWardrobe: () => client.get("/wardrobe/items"),
  analyzeItem: async (imageUri) => {
    const formData = new FormData();
    const filename = imageUri.split("/").pop();
    const match = /\.(\w+)$/.exec(filename);
    const type = match ? `image/${match[1]}` : `image`;

    formData.append("image", { uri: imageUri, name: filename, type });

    return client.post("/wardrobe/analyze", formData, {
      headers: { "Content-Type": "multipart/form-data" },
    });
  },
};
