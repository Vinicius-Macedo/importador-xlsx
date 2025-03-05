import {
  Sidebar,
  SidebarContent,
  SidebarGroup,
  SidebarGroupContent,
  SidebarGroupLabel,
  SidebarMenu,
  SidebarMenuButton,
  SidebarMenuItem,
  useSidebar,
} from "@/components/ui/sidebar";
import { Calendar, User, Inbox, Search, Settings, Users,ChartNoAxesCombined } from "lucide-react";
import { Link } from "react-router-dom";

const items = [
  {
    title: "Meu Perfil",
    url: "/usuario",
    icon: User,
  },
  {
    title: "Importar arquivo",
    url: "/usuario/importar-arquivo",
    icon: Inbox,
  },
  {
    title: "Clientes",
    url: "/usuario/clientes",
    icon: Users,
  },
  {
    title: "Recursos",
    url: "/usuario/recursos",
    icon: ChartNoAxesCombined,
  },
  {
    title: "Categorias",
    url: "/usuario/categorias",
    icon: Search,
  }
];

export function AppSidebar() {
  const { setOpenMobile } = useSidebar();

  return (
    <Sidebar>
      {/* <SidebarHeader /> */}
      <SidebarContent>
        <SidebarGroup />
        <SidebarGroupLabel>Importador</SidebarGroupLabel>
        <SidebarGroupContent>
          <SidebarMenu>
            {items.map((item) => (
              <SidebarMenuItem key={item.title}
                onClick={() => setOpenMobile(false)}
              >
                <SidebarMenuButton asChild>
                  <Link to={item.url}>
                    <item.icon />
                    <span>{item.title}</span>
                  </Link>
                </SidebarMenuButton>
              </SidebarMenuItem>
            ))}
          </SidebarMenu>
        </SidebarGroupContent>
        <SidebarGroup />
      </SidebarContent>
      {/* <SidebarFooter /> */}
    </Sidebar>
  );
}
