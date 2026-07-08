import { useEffect } from "react"
import { useLocation } from "react-router-dom"
import { PageViewService } from "../services/PageViewService"

const isAdminLoggedIn = (): boolean =>
  localStorage.getItem("profile") === "admin" &&
  !!localStorage.getItem("token")

export const usePageViewTracker = () => {
  const location = useLocation()

  useEffect(() => {
    if (isAdminLoggedIn()) return

    const path = location.pathname || "/"
    PageViewService.trackPageView(path).catch(() => {})
  }, [location.pathname])
}
