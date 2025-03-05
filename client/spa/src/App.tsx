import "./index.css";
import { Login } from "./pages/auth/Login";
import { Home } from "./pages/Home";
import { Register } from "./pages/auth/Register";
import { createBrowserRouter, RouterProvider } from "react-router-dom";
import { useQuery } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import { Profile } from "./pages/me/Profile";
import { ForgotPassword } from "./pages/auth/ForgotPassword";
import { ResetPassword } from "./pages/auth/ResetPassword";
import { SuccessEmail } from "./pages/auth/SuccessEmail";
import { SuccessResetPassword } from "./pages/auth/SuccessResetPassword";
import { Layout } from "./partials/Layout";
import { LateralLeftBarLayout } from "./partials/lateralLeftBarLayout";
import { ImportFile } from "./pages/me/ImportFile";
import { Customers } from "./pages/me/Customers";
import { Resources } from "./pages/me/Resources";
import { Categories } from "./pages/me/Categories";

export function App() {
  const { isLoading, isError } = useQuery({
    queryKey: ["pings"],
    queryFn: async () => {
      const { data } = await axiosInstance.get("/");
      return data;
    },
  });

  if (isLoading) {
    return null;
  }

  if (isError) {
    return <div>Internal server error, we working to fix this!</div>;
  }

  const router = createBrowserRouter(
    [
      {
        path: "/",
        element: (
          <Layout>
            <Home />
          </Layout>
        ),
      },
      {
        path: "/login",
        element: (
          <Layout>
            <Login />
          </Layout>
        ),
      },
      {
        path: "/register",
        element: (
          <Layout>
            <Register />
          </Layout>
        ),
      },
      {
        path: "/usuario",
        element: (
          <LateralLeftBarLayout>
            <Profile />
          </LateralLeftBarLayout>
        ),
      },
      {
        path: "/usuario/importar-arquivo",
        element: (
          <LateralLeftBarLayout>
            <ImportFile />
          </LateralLeftBarLayout>
        ),
      },
      {
        path: "/usuario/clientes",
        element: (
          <LateralLeftBarLayout>
            <Customers />
          </LateralLeftBarLayout>
        ),
      },
      {
        path: "/usuario/recursos",
        element: (
          <LateralLeftBarLayout>
            <Resources />
          </LateralLeftBarLayout>
        ),
      },
      {
        path: "/usuario/categorias",
        element: (
          <LateralLeftBarLayout>
            <Categories />
          </LateralLeftBarLayout>
        ),
      },

      {
        path: "/forgot-password",
        element: <ForgotPassword />,
      },
      {
        path: "/recover-password",
        element: <ResetPassword />,
      },
      {
        path: "/success-email",
        element: <SuccessEmail />,
      },
      {
        path: "/success-password",
        element: <SuccessResetPassword />,
      },
    ],
    {
      future: {
        v7_relativeSplatPath: true,
      },
    }
  );

  return (
    <>
      <RouterProvider router={router} />
      {/* <ReactQueryDevtools initialIsOpen={false} /> */}
    </>
  );
}
