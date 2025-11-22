import React, { useState, useRef } from "react";
import {
  View,
  Text,
  TouchableOpacity,
  StyleSheet,
  Image,
  Alert,
  ActivityIndicator,
} from "react-native";
import { CameraView, useCameraPermissions } from "expo-camera";
import { api } from "../api/client";

export default function CameraScreen({ navigation }) {
  const [permission, requestPermission] = useCameraPermissions();
  const [photo, setPhoto] = useState(null);
  const [loading, setLoading] = useState(false);
  const cameraRef = useRef(null);

  if (!permission) return <View />;
  if (!permission.granted) {
    return (
      <View style={styles.center}>
        <Text>Camera permission required</Text>
        <TouchableOpacity onPress={requestPermission} style={styles.permBtn}>
          <Text style={styles.text}>Grant</Text>
        </TouchableOpacity>
      </View>
    );
  }

  const takePicture = async () => {
    if (cameraRef.current) {
      const photoData = await cameraRef.current.takePictureAsync({
        quality: 0.5,
      });
      setPhoto(photoData.uri);
    }
  };

  const uploadPhoto = async () => {
    setLoading(true);
    try {
      await api.analyzeItem(photo);
      Alert.alert("Success", "Added to wardrobe!");
      setPhoto(null);
      navigation.navigate("Wardrobe");
    } catch (error) {
      Alert.alert("Error", "Ensure Go backend is running.");
    } finally {
      setLoading(false);
    }
  };

  if (photo)
    return (
      <View style={styles.container}>
        <Image source={{ uri: photo }} style={styles.preview} />
        {loading ? (
          <View style={styles.overlay}>
            <ActivityIndicator color="white" size="large" />
            <Text style={styles.text}>Analyzing...</Text>
          </View>
        ) : (
          <View style={styles.row}>
            <TouchableOpacity
              onPress={() => setPhoto(null)}
              style={[styles.btn, { backgroundColor: "#d9534f" }]}
            >
              <Text style={styles.text}>Retake</Text>
            </TouchableOpacity>
            <TouchableOpacity
              onPress={uploadPhoto}
              style={[styles.btn, { backgroundColor: "#4a5d23" }]}
            >
              <Text style={styles.text}>Analyze</Text>
            </TouchableOpacity>
          </View>
        )}
      </View>
    );

  return (
    <View style={styles.container}>
      <CameraView style={{ flex: 1 }} ref={cameraRef}>
        <View style={styles.cameraControls}>
          <TouchableOpacity style={styles.shutter} onPress={takePicture} />
        </View>
      </CameraView>
    </View>
  );
}

const styles = StyleSheet.create({
  container: { flex: 1 },
  center: { flex: 1, justifyContent: "center", alignItems: "center" },
  permBtn: {
    padding: 20,
    backgroundColor: "#4a5d23",
    borderRadius: 10,
    marginTop: 20,
  },
  cameraControls: {
    flex: 1,
    justifyContent: "flex-end",
    alignItems: "center",
    paddingBottom: 50,
  },
  shutter: {
    width: 70,
    height: 70,
    borderRadius: 35,
    backgroundColor: "white",
    borderWidth: 5,
    borderColor: "#ccc",
  },
  preview: { flex: 1, resizeMode: "contain", backgroundColor: "black" },
  row: {
    flexDirection: "row",
    justifyContent: "space-around",
    padding: 20,
    backgroundColor: "black",
  },
  btn: { padding: 15, borderRadius: 8, minWidth: 100, alignItems: "center" },
  text: { color: "white", fontWeight: "bold" },
  overlay: {
    ...StyleSheet.absoluteFillObject,
    backgroundColor: "rgba(0,0,0,0.7)",
    justifyContent: "center",
    alignItems: "center",
  },
});
