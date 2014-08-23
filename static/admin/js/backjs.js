function previewUpload(id) {
	$.get('/api/common/upload?id=' + id, function(r) {
		if (r.ok) {
			window.open('/'+r.filepath)
		}
	})
}
function putFormUpload(form) {
	$.ajax({
			 type: "PUT",
			 url: "/api/common/upload",
			 data: form.serialize(),
			 contentType:"multipart/form-data; boundary=----WebKitFormBoundary8qkcpEF64rHWHlkd",
			 success: function(msg){
				  console.log(msg)
			 }
	});
}

function deletePost(url, infoelem) {
	$.ajax({
		url: url,
		type: 'DELETE',
		success: function(r) {
			handleResp(r, infoelem)					
		}
	})
}

function putPost(url, form, infoelem) {
	$.ajax({
		url: url,
		type: 'PUT',
		data: form.serialize(),
		success: function(r) {
			handleResp(r, infoelem)					
		}
	})
}

function postPost(url, form, infoelem) {
	$.ajax({
		url: url,
		type: 'POST',
		data: form.serialize(),
		success: function(r) {
			handleResp(r, infoelem)					
		}
	})
}

function handleResp(re, infoelem) {
	if (re.ok) {
		location.reload()	
	} else {
		infoelem.after('<div class="alert alert-danger" role="alert">'+r.errmsg+'</div>')
	}
}