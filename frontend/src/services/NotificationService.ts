import ApiPublica from "../providers/ApiPublica";

export const NotificationService = {
  sendWelcome: (userId: number) =>
    ApiPublica.post("/notifications/welcome", { userId }),
};
