import { useState } from "react";
import { useQuery } from "@tanstack/react-query";
import { axiosInstance } from "@/services/axiosInstance";
import {
  Table,
  TableBody,
  TableCaption,
  TableCell,
  TableFooter,
  TableHead,
  TableHeader,
  TableRow,
} from "@/components/ui/table";
import { Input } from "@/components/ui/input";
import { Label } from "@/components/ui/label";

interface CategoriesProps {}

export function Categories(props: CategoriesProps) {
  const [searchTerm, setSearchTerm] = useState("");
  const { isLoading, isError, data } = useQuery({
    queryKey: ["categories"],
    queryFn: async () => {
      const response = await axiosInstance.get("/categories");
      return response.data; // Return the data directly
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error loading data</div>;
  }

  const filteredData = data ? data.filter((category: any) =>
    category.name.toLowerCase().includes(searchTerm.toLowerCase())
  ) : [];

  return (
    <div className="p-6 w-full responsive-table">
      <div className="mb-4">
        <Label htmlFor="search">Pesquisar Categoria</Label>
        <Input
          id="search"
          type="text"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
          placeholder="Digite para pesquisar..."
          className="mt-2"
        />
      </div>
      <div className="overflow-x-auto">
        {filteredData.length > 0 ? (
          <Table className="w-full">
            <TableCaption>Lista de Categorias</TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead>Nome</TableHead>
                <TableHead>Categoria</TableHead>
                <TableHead>Sub Categoria</TableHead>
                <TableHead>Tipo</TableHead>
                <TableHead>Unidade</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredData.map((category: any) => (
                <TableRow key={category.meters_key}>
                  <TableCell>{category.name}</TableCell>
                  <TableCell>{category.category}</TableCell>
                  <TableCell>{category.sub_category}</TableCell>
                  <TableCell>{category.type}</TableCell>
                  <TableCell>{category.unit}</TableCell>
                </TableRow>
              ))}
            </TableBody>
            <TableFooter>
              <TableRow>
                <TableCell colSpan={4}>Total</TableCell>
                <TableCell className="text-right">
                  {filteredData.length}
                </TableCell>
              </TableRow>
            </TableFooter>
          </Table>
        ) : (
          <div className="text-center py-4">Nenhuma categoria encontrada.</div>
        )}
      </div>
    </div>
  );
}