namespace go advertiseproject.advertise
include "base.thrift"
struct Advertise {
    1: i64 id
    2: string description
    3: string name
    4: i64 stock
}

struct GetAdReq {
    1: required i64 id
    2: string name
}

struct GetAdRes {
    1: Advertise ad

    255: base.BaseRes baseRes
}

struct AddAdReq {
    1: Advertise ad
}

struct AddAdRes {

    255: base.BaseRes baseRes
}

struct DeleteAdReq {
    1: required i64 id
    2: string name
}

struct DeleteAdRes {
    255: base.BaseRes res
}

struct UpdateAdReq {
    1: Advertise ad
}

struct UpdateAdRes {
    255: base.BaseRes res
}


service AdService {
    GetAdRes GetAd(1: GetAdReq req)
    AddAdRes AddAd(1: AddAdReq req)
    DeleteAdRes DeleteAd(1: DeleteAdReq req)
    UpdateAdRes UpdateAd(1: UpdateAdReq req)
}
