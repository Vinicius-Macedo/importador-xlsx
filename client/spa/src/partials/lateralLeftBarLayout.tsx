import { useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { useAuth } from "@/hooks/useAuth";
import { SidebarProvider, SidebarTrigger } from "@/components/ui/sidebar";
import { AppSidebar } from "@/components/app-sidebar";

interface lateralLeftBarLayoutProps {
  children: React.ReactNode;
}
export function LateralLeftBarLayout(props: lateralLeftBarLayoutProps) {
  const { isLoading, isAuthenticated } = useAuth();
  const navigate = useNavigate();
  const authUrls = ["/usuario", "/usuario/categorias"];

  useEffect(() => {
    if (
      isAuthenticated === false &&
      authUrls.includes(window.location.pathname)
    ) {
      navigate("/login");
    }
  }, [isAuthenticated]);

  if (isAuthenticated === null) {
    return null;
  }

  if (isLoading) {
    return null;
  }

  return (
    <>
      <SidebarProvider>
        <AppSidebar />
        <div className="flex flex-col flex-1">
          <header className="bg-gray-100 h-12 flex justify-between items-center">
            <SidebarTrigger className={"hover:bg-transparent"} />
            <p className="px-8">Bem-vindo</p>
          </header>
          <main>{props.children}</main>
        </div>
      </SidebarProvider>
    </>
  );
}
