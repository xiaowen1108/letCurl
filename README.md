# letCurl
go版本的http版本简单封装


curl := letCurl.NewCurl()


content, err:= curl.SetUrl("http://www.sasa.com/product-ajax_get_product_info.html").SetForm("product_id_str", "201,5547").Post()


//content是string类型  err是error类型


if err == nil {


		fmt.Println(content)
		
		
} else {


		fmt.Println(err)
		
		
}
