package biz

import (
    "context"
    v1 "github.com/BingguWang/bingfood-client-micro/api/cart/service/v1/pbgo/v1"
    v2 "github.com/BingguWang/bingfood-client-micro/api/prod/service/v1/pbgo/v1"
    "github.com/BingguWang/bingfood-client-micro/app/cart/service/internal/data/entity"
    "github.com/BingguWang/bingfood-client-micro/app/cart/service/internal/utils"
    "github.com/go-kratos/kratos/v2/log"
    "github.com/jinzhu/copier"
)

type CartRepo interface {
    AddCartRepo(context.Context, *entity.Cart) error
    GetCart(context.Context, *entity.Cart, int, int) ([]*entity.Cart, int64, error)
    GetCartByIds(context.Context, []uint64, int, int) ([]*entity.Cart, int64, error)
    //Update(context.Context, *entity.Cart) error
    //FindByID(context.Context, int64) (*entity.Cart, error)
    //ListAll(context.Context) ([]*entity.Cart, error)
    //GetUserByCond(ctx context.Context, user *entity.Cart) ([]*entity.Cart, error)
}

type CartUseCase struct {
    repo CartRepo
    usc  v2.ProdServiceClient
    log  *log.Helper
}

func NewCartUseCase(repo CartRepo, usc v2.ProdServiceClient, logger log.Logger) *CartUseCase {
    return &CartUseCase{repo: repo, usc: usc, log: log.NewHelper(logger)}
}

func (uc *CartUseCase) AddCartItem(ctx context.Context, req *v1.AddCartItemRequest) error {
    uc.log.WithContext(ctx).Infof("AddCartItem args: %v", utils.ToJsonString(req))

    if req.ProdNums <= 0 {
        // TODO _ = bs.DeleteBasketRecord(reqParam)
        return v1.ErrorInvalidArgument("商品已移出购物车")
    }
    cartRow := &entity.Cart{
        CartId:   req.CartId,
        ShopId:   req.ShopId,
        SkuId:    req.SkuId,
        UserId:   req.UserId,
        ProdNums: req.ProdNums,
    }
    if err := uc.repo.AddCartRepo(ctx, cartRow); err != nil {
        return v1.ErrorInternal("新增失败, internal error : ", err.Error())
    }

    return nil
}

func (uc *CartUseCase) GetCartHandler(ctx context.Context, req *v1.GetCartByCondRequest) ([]*entity.Cart, int64, error) {
    uc.log.WithContext(ctx).Infof("request args are : %v", utils.ToJsonString(req))

    if req.PageInfo == nil || req.CartCond == nil {
        return nil, 0, v1.ErrorInvalidArgument("错误, 请检查是否按要求传递参数")
    }

    var cart entity.Cart
    _ = copier.CopyWithOption(&cart, req.CartCond, copier.Option{
        IgnoreEmpty: false,
        DeepCopy:    true,
    })
    limit := (int)(req.PageInfo.PageSize)
    offset := (int)(req.PageInfo.Page)

    ret, total, err := uc.repo.GetCart(ctx, &cart, limit, offset)
    if err != nil {
        return nil, 0, v1.ErrorInternal("获取购物车失败, internal error : %v", err.Error())
    }

    if e := uc.FillSkuInfo(ctx, &ret); e != nil {
        return nil, 0, v1.ErrorInternal("获取sku信息失败, internal error : %v", e.Error())
    }
    return ret, total, nil
}

func (uc *CartUseCase) GetCartByIdsHandler(ctx context.Context, req *v1.GetCartByCartIdsRequest) ([]*entity.Cart, int64, error) {
    uc.log.WithContext(ctx).Infof("request args are : %v", utils.ToJsonString(req))

    if req.PageInfo == nil {
        return nil, 0, v1.ErrorInvalidArgument("错误, 请检查是否按要求传递参数")
    }
    limit := (int)(req.PageInfo.PageSize)
    offset := (int)(req.PageInfo.Page)
    ret, total, err := uc.repo.GetCartByIds(ctx, req.Ids, limit, offset)

    if err != nil {
        return nil, 0, v1.ErrorInternal("获取购物车失败, internal error : %v", err.Error())
    }

    if err := uc.FillSkuInfo(ctx, &ret); err != nil {
        return nil, 0, v1.ErrorInternal("获取sku信息失败, internal error : %v", err.Error())
    }
    return ret, total, nil
}

func (uc *CartUseCase) FillSkuInfo(ctx context.Context, ret *[]*entity.Cart) error {
    for _, ca := range *ret {
        rt, err := uc.usc.GetSkuByCond(ctx, &v2.GetSkuByCondRequest{
            SkuCond: &v2.Sku{SkuId: ca.SkuId},
        })
        log.Infof("调用服务bingfood.prod.service/GetSkuByCond, 得到结果: %v ", utils.ToJsonString(rt))
        if err != nil {
            return err
        }
        if rt.Data.List == nil {
            continue
        }
        copier.Copy(&ca.Sku, rt.Data.List[0])
    }
    return nil
}