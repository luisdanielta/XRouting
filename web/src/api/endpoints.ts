export const Endpoints = {
  auth: {
    login: "/auth/login",
    signup: "/auth/signup",
    logout: "/auth/logout",
  },
  chat: {
    messages: "/msg",
    websocket: "ws://localhost:3000/ws", // WebSocket endpoint for real-time communication
  },
};
