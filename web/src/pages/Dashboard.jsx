import React, { useEffect } from "react";
import { useNavigate } from "react-router-dom";

const Dashboard = () => {
  const navigate = useNavigate();

  useEffect(() => {
    const params = new URLSearchParams(window.location.search);
    const token = params.get("token");

    if (token) {
      localStorage.setItem("token", token); // Lưu token vào localStorage
      navigate("/dashboard"); // Xóa query params khỏi URL
    }
  }, [navigate]);

  return (
    <div>
      <h1>Welcome to your Dashboard</h1>
      <p>This is a protected route. You are logged in!</p>
    </div>
  );
};

export default Dashboard;