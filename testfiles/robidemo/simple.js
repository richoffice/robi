; (function (input) {
    src = input[0];
    rfs = robi.Import("sample_def.json", src);
    visitors = rfs.Get("visitors");
    log(visitors);

    var addFunc = function(row){
        return row["name"]+"-xxx";
    };
    visitors.Add("fullname",addFunc);
    log(rfs);
    return src;
})(input)