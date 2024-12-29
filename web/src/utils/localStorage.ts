// base64加密存储token
const setToken = (token) => {
    // 将token字段加密存储
    localStorage.setItem(window.btoa("token"), window.btoa(token));
}
// 获取base64解密后的token
const getToken = () => {
    return window.atob(localStorage.getItem(window.btoa("token")));
}
// 删除token
const removeToken = () => {
    localStorage.removeItem(window.btoa("token"));
}

// 存储user obj
const setUser = (user) => {
    localStorage.setItem(window.btoa("user"), window.btoa(JSON.stringify(user)));
}
// 获取user obj
const getUser = () => {
    return JSON.parse(window.atob(localStorage.getItem(window.btoa("user"))));
}
// 删除user obj
const removeUser = () => {
    localStorage.removeItem(window.btoa("user"));
}
export { setToken,
    getToken,
    removeToken,
    setUser,
    getUser,
    removeUser
};
