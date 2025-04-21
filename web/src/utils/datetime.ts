import { parseDateTime } from "@internationalized/date";

// 获取当前日期时间
// 返回格式为 2023-07-01T09:00:00
export const getCurrentDateTime = () => {
  const Dates = new Date();
  // 年份
  const Year = Dates.getFullYear();
  // 月份下标是0-11
  const Months =
    Dates.getMonth() + 1 < 10
      ? "0" + (Dates.getMonth() + 1)
      : Dates.getMonth() + 1;
  // 具体的天数
  const Day = Dates.getDate() < 10 ? "0" + Dates.getDate() : Dates.getDate();
  // 小时
  const Hours =
    Dates.getHours() < 10 ? "0" + Dates.getHours() : Dates.getHours();
  // 分钟
  const Minutes =
    Dates.getMinutes() < 10 ? "0" + Dates.getMinutes() : Dates.getMinutes();
  // 秒
  const Seconds =
    Dates.getSeconds() < 10 ? "0" + Dates.getSeconds() : Dates.getSeconds();

  // 返回数据格式
  return parseDateTime(
    `${Year}-${Months}-${Day}T${Hours}:${Minutes}:${Seconds}`,
  );
};

// 格式化日期函数
// 将 2023-07-01T09:00:00 格式化为 07-01 09:00
export const formatDate = (dateString: string | undefined) => {
  if (!dateString) return "";
  const date = new Date(dateString);
  const options: Intl.DateTimeFormatOptions = {
    month: "2-digit",
    day: "2-digit",
    hour: "2-digit",
    minute: "2-digit",
    hour12: false,
  };

  return new Intl.DateTimeFormat("zh-CN", options).format(date);
};
