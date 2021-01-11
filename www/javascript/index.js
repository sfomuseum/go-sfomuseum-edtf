function to_edtf_string(date){

    var on_success = function(rsp){
	console.log("SUCCESS", rsp);
    };

    var on_error = function(err){
	console.log("ERROR", err);
    };

    console.log("TO EDTF STRING", date);
    sfomuseum.date.toEDTFString(date, on_success, on_error);
}
