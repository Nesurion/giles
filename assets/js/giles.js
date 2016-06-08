function archivMovieRequest(title, path){
	$.ajax({
	    url: 'api/archiv/movie',
	    type: 'POST',
	    data: {title: title, path: path}
	});
}

function archivShowRequest(title, path){
	$.ajax({
	    url: 'api/archiv/show',
	    type: 'POST',
	    data: {title: title, path: path}
	});
}

function deleteRequest(title, path){
	$.ajax({
	    url: 'api/delete',
	    type: 'POST',
	    data: {title: title, path: path}
	});
}
