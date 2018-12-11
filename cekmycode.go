
// Package Controller
func CreateSuratJalan(c *gin.Context) {
	data := c.PostFormArray("data")
	fmt.Println(data)

	var surat model.SuratJalanStruct
	err := c.Bind(&surat)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(surat)
}
//==========================


// Package Model
type DataJalan struct {
	Sku    []string `json:"sku"`
	Ukuran []string `json:"ukuran"`
	Warna  []string `json:"warna"`
	Qty    []string `json:"qty"`
}

type SuratJalanStruct struct {
	Data DataJalan `json:"data"`
}
//========================

// JS 
 $("#buatsurat").click(function(){
    console.log(Detail);
    $.ajax({
      url : "/dashboard/fulfillment/create/suratjalan",
      dataType: "json",
      type : "POST",
      data : {
        "data" : Detail,
      },
    });
  });
//========================
