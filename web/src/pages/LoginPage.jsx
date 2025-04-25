import React from "react";

const LoginPage = () => {
  const handleLogin = () => {
    window.location.href = "http://localhost:8080/v1/auth/google";
  };

  return (
    <div style={{ textAlign: "center", marginTop: "50px" }}>
      <h1>Welcome to Task App</h1>
      <button onClick={handleLogin}>Login with Google</button>
    </div>
  );
};

export default LoginPage;