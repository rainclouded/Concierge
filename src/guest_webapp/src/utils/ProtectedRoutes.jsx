import { Navigate } from "react-router-dom";

//See https://www.youtube.com/watch?v=pyfwQUc5Ssk
export const ProtectedRoute = ( {canAccessFunc, element, fallbackPath} ) => {
    return canAccessFunc() ? element : <Navigate to={fallbackPath} />;
}