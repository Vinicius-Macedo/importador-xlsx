import { Footer } from "./Footer";
import { Header } from "./Header";

interface LayoutProps {
  children: React.ReactNode;
}
export function Layout(props : LayoutProps){
    return(
        <>
        <Header/>
          <main className="flex flex-col flex-auto pt-[72px]">
            {props.children}
          </main>
        <Footer/>
        </>
    )
}