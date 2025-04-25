import React from "react";
import { Navigate } from "react-router-dom";

const ProtectedRoute = ({ children }) => {
  const token = localStorage.getItem("token"); // Lấy token từ localStorage

  if (!token) {
    return <Navigate to="/" />; // Chuyển hướng về trang login nếu chưa đăng nhập
  }

  return children;
};

export default ProtectedRoute;