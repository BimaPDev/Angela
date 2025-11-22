import React, { useState, useCallback } from "react";
import {
  View,
  Text,
  FlatList,
  Image,
  StyleSheet,
  ActivityIndicator,
} from "react-native";
import { useFocusEffect } from "@react-navigation/native";
import { api } from "../api/client";

export default function WardrobeScreen() {
  const [items, setItems] = useState([]);
  const [loading, setLoading] = useState(true);

  useFocusEffect(
    useCallback(() => {
      loadData();
    }, [])
  );

  const loadData = async () => {
    try {
      const response = await api.getWardrobe();
      setItems(response.data);
    } catch (error) {
      // Fail silently or log error if backend is down
      console.log("Backend not reachable yet");
    } finally {
      setLoading(false);
    }
  };

  const renderItem = ({ item }) => (
    <View style={styles.card}>
      <Image source={{ uri: item.image_url }} style={styles.image} />
      <View style={styles.textContainer}>
        <Text style={styles.category}>{item.category}</Text>
        <Text style={styles.details}>
          {item.color} • {item.style}
        </Text>
      </View>
    </View>
  );

  return (
    <View style={styles.container}>
      <Text style={styles.header}>My Wardrobe</Text>
      {loading ? (
        <ActivityIndicator />
      ) : (
        <FlatList
          data={items}
          renderItem={renderItem}
          keyExtractor={(item) =>
            item.ID ? item.ID.toString() : Math.random().toString()
          }
          numColumns={2}
        />
      )}
    </View>
  );
}

const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#F5F5DC",
    paddingTop: 60,
    paddingHorizontal: 10,
  },
  header: {
    fontSize: 28,
    fontWeight: "bold",
    color: "#4a5d23",
    marginBottom: 15,
  },
  card: {
    flex: 1,
    margin: 5,
    backgroundColor: "white",
    borderRadius: 12,
    padding: 10,
  },
  image: {
    width: "100%",
    height: 120,
    borderRadius: 8,
    backgroundColor: "#eee",
  },
  textContainer: { marginTop: 8 },
  category: { fontWeight: "bold" },
  details: { fontSize: 12, color: "#666" },
});
