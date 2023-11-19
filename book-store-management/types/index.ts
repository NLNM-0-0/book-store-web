import { IconType } from "react-icons";

export type Book = {
  id: string;
  name: string;
  nxb: string;
  quantity: number;
  price: number;
  status: boolean;
  category: string;
};
export type ImportNote = {
  id: string;
  supplierId: string;
  totalPrice: number;
  status: StatusNote;
  createBy: string;
  closeBy?: string;
  createAt: Date;
  closeAt?: Date;
};

export type ImportDetail = {
  book: Book;
  idNote: string;
  quantity: number;
  price: number;
};

export type Supplier = {
  id: string;
  name: string;
};

export enum StatusNote {
  Inprogress = "Đang xử lý",
  Done = "Đã nhập",
  Cancel = "Đã huỷ",
}
export enum StatusActive {
  Active = "Đang giao dịch",
  InActive = "Ngừng giao dịch",
}
export type Category = {
  id: string;
  name: string;
};

export interface CategoryListProps {
  category: string;
  setCategory: (category: string) => void;
  canAdd?: boolean;
  readonly?: boolean;
}
export interface SupplierListProps {
  supplier: string;
  setSupplier: (supplier: string) => void;
  canAdd?: boolean;
}
export interface BookListProps {
  book: Partial<Book>;
  setBook: (book: Partial<Book>) => void;
  isNew: boolean;
  setIsNew: (isNew: boolean) => void;
}

export type SidebarItem = {
  title: string;
  href: string;
  icon?: IconType;
  submenu?: boolean;
  subMenuItems?: SidebarItem[];
};