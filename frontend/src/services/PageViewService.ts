import Api from "../providers/Api"
import ApiPublica from "../providers/ApiPublica"

export interface IPageView {
  id: number
  path: string
  viewDate: string
  count: number
}

const trackPageView = (path: string) =>
  ApiPublica.post("/page-view", { path })

const getPageViews = () => Api.get<IPageView[]>("/page-views")

export const PageViewService = {
  trackPageView,
  getPageViews,
}
