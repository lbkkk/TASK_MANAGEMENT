import React from "react";

const GoogleLogin = () => {
  const handleLogin = () => {
    window.location.href = "http://localhost:8080/v1/auth/google";
  };

  return <button onClick={handleLogin}>Login with Google</button>;
};

export default GoogleLogin;