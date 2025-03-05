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
import { Button } from "@/components/ui/button";

interface ResourcesProps {}

export function Resources(props: ResourcesProps) {
  const [searchTerm, setSearchTerm] = useState("");
  const { isLoading, isError, data } = useQuery({
    queryKey: ["resources"],
    queryFn: async () => {
      const response = await axiosInstance.get("/resources");
      return response.data; // Return the data directly
    },
  });

  if (isLoading) {
    return <div>Loading...</div>;
  }

  if (isError) {
    return <div>Error loading data</div>;
  }

  const filteredData = data
    ? data.filter((resource: any) =>
        resource.resource_group.toLowerCase().includes(searchTerm.toLowerCase())
      )
    : [];

  return (
    <div className="p-6 w-full responsive-table">
      <div className="mb-4">
        <Label htmlFor="search">Pesquisar Grupo de Recursos</Label>
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
          <Table className="min-w-full">
            <TableCaption>Lista de Recursos</TableCaption>
            <TableHeader>
              <TableRow>
                <TableHead>Grupo de Recursos</TableHead>
                <TableHead>Serviço Consumido</TableHead>
                <TableHead>Localização</TableHead>
                <TableHead>URI</TableHead>
              </TableRow>
            </TableHeader>
            <TableBody>
              {filteredData.map((resource: any) => (
                <TableRow key={resource.id}>
                  <TableCell>{resource.resource_group}</TableCell>
                  <TableCell>{resource.consumed_service}</TableCell>
                  <TableCell>{resource.location}</TableCell>
                  <TableCell>{resource.uri}</TableCell>
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
          <div className="text-center py-4">Nenhum grupo de recurso encontrado.</div>
        )}
      </div>
    </div>
  );
}
