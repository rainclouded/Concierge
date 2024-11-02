import { getSessionkey } from "./auth";

const SESSION_KEY = "X-API-Key"

export const fetchWithAuth = async (url, options = {}) => {
    options.headers = {
      ...options.headers,
      [SESSION_KEY]: getSessionkey(),
    };

    return fetch(url, options)
}