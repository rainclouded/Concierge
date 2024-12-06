import { jwtDecode } from "jwt-decode";

const SESSION_KEY = "sessionKey";
let cachedAccountId = null;

export const setSessionKey = (key) => {
  localStorage.setItem(SESSION_KEY, key);
  try {
    const decoded = jwtDecode(key);
    cachedAccountId = decoded?.accountId; // Store accountId in cache
  } catch {
    cachedAccountId = null;
  }
};

export const getSessionkey = () => {
  return localStorage.getItem(SESSION_KEY);
};

export const removeSessionKey = () => {
  localStorage.removeItem(SESSION_KEY);
};

export const isAuthenticated = () => {
  const sessionKey = getSessionkey();
  if (!sessionKey || sessionKey == "") {
    return false;
  }

  try {
    const decoded = jwtDecode(sessionKey);
    console.log(decoded);
    const remainingTime = (decoded.exp ?? 0) - Math.floor(Date.now() / 1000);
    console.log(
      `Key remaining lifespan: ${Math.floor(remainingTime / 3600)}:${Math.floor(
        (remainingTime % 3600) / 60
      )}:${Math.floor(remainingTime % 60)}`
    );
    return remainingTime > 0;
  } catch (error) {
    console.log(error);
    return false;
  }
};

// Utility to get the accountId from the cached value or by decoding the token
export const getAccountId = () => {
  if (cachedAccountId) {
    return cachedAccountId;
  }

  const sessionKey = getSessionkey();
  if (!sessionKey) {
    return null;
  }

  try {
    const decoded = jwtDecode(sessionKey);
    cachedAccountId = decoded?.accountId;
    return cachedAccountId;
  } catch {
    return null;
  }
};
