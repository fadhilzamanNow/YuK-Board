export default defineNuxtRouteMiddleware((to) => {
  if (import.meta.client) {
    const token = localStorage.getItem("authToken");
    
    const publicRoutes = ["/", "/login", "/register"];
    const isPublicRoute = publicRoutes.includes(to.path);
    
    if (!token && !isPublicRoute) {
      return navigateTo("/login");
    }
    
    if (token && (to.path === "/login" || to.path === "/register" || to.path === "/")) {
      return navigateTo("/home");
    }
  }
});
