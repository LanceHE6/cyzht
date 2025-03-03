import { addToast } from "@heroui/react";

export class ToastProvider {
  public default(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "default",
    });
  }

  public danger(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "danger",
    });
  }
  public primary(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "primary",
    });
  }
  public success(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "success",
    });
  }
  public foreground(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "foreground",
    });
  }
  public secondary(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "secondary",
    });
  }
  public warning(title: string, message: string | null) {
    addToast({
      title: title,
      description: message ? message : undefined,
      color: "warning",
    });
  }
}
