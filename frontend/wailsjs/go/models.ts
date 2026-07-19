export namespace models {
	
	export class Product {
	    id: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    category_id: string;
	    category?: Category;
	    name: string;
	    sku: string;
	    barcode: string;
	    price: number;
	    stock: number;
	    is_active: boolean;
	
	    static createFrom(source: any = {}) {
	        return new Product(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.category_id = source["category_id"];
	        this.category = this.convertValues(source["category"], Category);
	        this.name = source["name"];
	        this.sku = source["sku"];
	        this.barcode = source["barcode"];
	        this.price = source["price"];
	        this.stock = source["stock"];
	        this.is_active = source["is_active"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Category {
	    id: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    name: string;
	    products?: Product[];
	
	    static createFrom(source: any = {}) {
	        return new Category(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.name = source["name"];
	        this.products = this.convertValues(source["products"], Product);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class SaleItem {
	    id: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    sale_id: string;
	    product_id: string;
	    product?: Product;
	    name_snapshot: string;
	    price_snapshot: number;
	    qty: number;
	    line_total: number;
	
	    static createFrom(source: any = {}) {
	        return new SaleItem(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.sale_id = source["sale_id"];
	        this.product_id = source["product_id"];
	        this.product = this.convertValues(source["product"], Product);
	        this.name_snapshot = source["name_snapshot"];
	        this.price_snapshot = source["price_snapshot"];
	        this.qty = source["qty"];
	        this.line_total = source["line_total"];
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Sale {
	    id: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    invoice_no: string;
	    customer_id?: string;
	    customer?: Customer;
	    subtotal: number;
	    tax: number;
	    discount: number;
	    total: number;
	    paid: number;
	    change: number;
	    payment_method: string;
	    items?: SaleItem[];
	
	    static createFrom(source: any = {}) {
	        return new Sale(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.invoice_no = source["invoice_no"];
	        this.customer_id = source["customer_id"];
	        this.customer = this.convertValues(source["customer"], Customer);
	        this.subtotal = source["subtotal"];
	        this.tax = source["tax"];
	        this.discount = source["discount"];
	        this.total = source["total"];
	        this.paid = source["paid"];
	        this.change = source["change"];
	        this.payment_method = source["payment_method"];
	        this.items = this.convertValues(source["items"], SaleItem);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	export class Customer {
	    id: string;
	    // Go type: time
	    created_at: any;
	    // Go type: time
	    updated_at: any;
	    name: string;
	    phone: string;
	    sales?: Sale[];
	
	    static createFrom(source: any = {}) {
	        return new Customer(source);
	    }
	
	    constructor(source: any = {}) {
	        if ('string' === typeof source) source = JSON.parse(source);
	        this.id = source["id"];
	        this.created_at = this.convertValues(source["created_at"], null);
	        this.updated_at = this.convertValues(source["updated_at"], null);
	        this.name = source["name"];
	        this.phone = source["phone"];
	        this.sales = this.convertValues(source["sales"], Sale);
	    }
	
		convertValues(a: any, classs: any, asMap: boolean = false): any {
		    if (!a) {
		        return a;
		    }
		    if (a.slice && a.map) {
		        return (a as any[]).map(elem => this.convertValues(elem, classs));
		    } else if ("object" === typeof a) {
		        if (asMap) {
		            for (const key of Object.keys(a)) {
		                a[key] = new classs(a[key]);
		            }
		            return a;
		        }
		        return new classs(a);
		    }
		    return a;
		}
	}
	
	

}

