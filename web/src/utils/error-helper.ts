// getErrorMessage 获取错误信息
export function getErrorMessage(error: unknown): string {
  if (error instanceof Error) {
    return error.message;
  } else if (
    typeof error === "object" &&
    error !== null &&
    "response" in error &&
    typeof error.response === "object" &&
    error.response !== null &&
    "data" in error.response &&
    typeof error.response.data === "object" &&
    error.response.data !== null &&
    "msg" in error.response.data
  ) {
    return (error.response.data.msg as string) || "未知错误";
  }

  return "未知错误";
}
