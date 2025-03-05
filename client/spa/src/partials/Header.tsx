import { useAuth } from "@/hooks/useAuth";
import { Link, useNavigate } from "react-router-dom";
import { axiosInstance } from "@/services/axiosInstance";
import { useQueryClient } from "@tanstack/react-query";
import { useState } from "react";

export function Header() {
  const { isAuthenticated } = useAuth();
  const [isMenuOpen, setIsMenuOpen] = useState(false);

  const navigate = useNavigate();
  const queryClient = useQueryClient();

  async function handleLogout() {
    await axiosInstance.post("/logout");
    queryClient.removeQueries({ queryKey: ["repoData"] });
    navigate("/");
  }

  return (
    <>
      <div
        className={
          "fixed left-0 top-0 w-screen h-screen z-40 lg:hidden" +
          (isMenuOpen ? " block" : " hidden")
        }
        onClick={() => setIsMenuOpen(false)}
      ></div>
      <nav
        className="fixed left-0 top-0 w-full z-50 bg-slate-50"
        role="navigation"
        aria-label="main navigation"
      >
        <div className="container mx-auto px-2 sm:px-6 lg:px-8 flex items-center w-full">
          <div className="flex items-center justify-between flex-wrap w-full max-w-7xl mx-auto px-2 sm:px-6 lg:px-8">
            <Link to={"/"} className="shrink">
              <h1 className="font-bold">LOGO</h1>
            </Link>
            <div
              className="block lg:hidden"
              onClick={() => setIsMenuOpen(!isMenuOpen)}
            >
              <button
                id="menu-toggle"
                className="flex items-center px-3 py-6 rounded text-black-500 hover:text-black-400"
              >
                <svg
                  id="menu-open"
                  className="fill-current h-5 w-5 block"
                  viewBox="0 0 20 20"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path d="M0 3h20v2H0V3zm0 6h20v2H0V9zm0 6h20v2H0v-2z"></path>
                </svg>
                <svg
                  id="menu-close"
                  className="fill-current h-5 w-5 hidden"
                  viewBox="0 0 20 20"
                  xmlns="http://www.w3.org/2000/svg"
                >
                  <path d="M10 8.586L2.929 1.515 1.515 2.929 8.586 10l-7.071 7.071 1.414 1.414L10 11.414l7.071 7.071 1.414-1.414L11.414 10l7.071-7.071-1.414-1.414L10 8.586z"></path>
                </svg>
              </button>
            </div>
            <div
              id="menu"
              className={
                "w-full block flex-grow lg:flex lg:items-center lg:w-auto lg:justify-end" +
                (isMenuOpen ? " block" : " hidden")
              }
            >
              {isAuthenticated ? (
                <>
                  <Link
                    to={"/usuario"}
                    onClick={() => setIsMenuOpen(false)}
                    className="block lg:inline-block p-4 lg:py-6 text-center text-[16px] text-white-200"
                  >
                    System
                  </Link>

                  <p
                    className="block lg:inline-block p-4 lg:py-6 text-center text-[16px] text-white-200 cursor-pointer"
                    onClick={() => {
                      handleLogout();
                      setIsMenuOpen(false);
                    }}
                  >
                    Logout
                  </p>
                </>
              ) : (
                <>
                  <Link
                    onClick={() => setIsMenuOpen(false)}
                    to={"/Login"}
                    className="block lg:inline-block p-4 lg:py-6 text-center text-[16px] text-white-200"
                  >
                    Login
                  </Link>
                  <Link
                    onClick={() => setIsMenuOpen(false)}
                    to={"/register"}
                    className="block lg:inline-block p-4 lg:py-6 text-center text-[16px] text-white-200"
                  >
                    Cadastre-se
                  </Link>
                </>
              )}
            </div>
          </div>
        </div>
      </nav>
    </>
  );
}
