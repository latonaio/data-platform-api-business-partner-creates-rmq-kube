# data-platform-api-business-partner-creates-rmq-kube
data-platform-api-business-partner-creates-rmq-kube は、周辺システム　を データ連携基盤 と統合することを目的に、API でビジネスパートナデータを登録/更新するマイクロサービスです。

* https://xxx.xxx.io/api/API_BUSINESS_PARTNER_SRV/creates/
* https://xxx.xxx.io/api/API_BUSINESS_PARTNER_SRV/updates/

## 動作環境
data-platform-api-business-partner-creates-rmq-kube の動作環境は、次の通りです。  
・ OS: LinuxOS （必須）  
・ CPU: ARM/AMD/Intel（いずれか必須）  

## 本レポジトリ が 対応する API サービス
data-platform-api-business-partner-creates-rmq-kube が対応する APIサービス は、次のものです。

* APIサービス URL: https://xxx.xxx.io/api/API_BUSINESS_PARTNER_SRV/creates/
* APIサービス URL: https://xxx.xxx.io/api/API_BUSINESS_PARTNER_SRV/updates/

## 本レポジトリ に 含まれる API名
data-platform-api-business-partner-creates-rmq-kube には、次の API をコールするためのリソースが含まれています。  

* A_General（ビジネスパートナ - 一般）
* A_GeneralDoc（ビジネスパートナ - 一般文書）
* A_Role（ビジネスパートナ - ロール）
* A_Person（ビジネスパートナ - 個人）
* A_Address（ビジネスパートナ - 住所）
* A_SNS（ビジネスパートナ - SNS）
* A_GPS（ビジネスパートナ - GPS）
* A_Rank（ビジネスパートナ - ランク）
* A_PersonOrganization（ビジネスパートナ - 個人組織）
* A_PersonMobilePhoneAuth（ビジネスパートナ - 個人携帯電話認証）
* A_PersonGoogleAccountAuth（ビジネスパートナ - 個人Googleアカウント認証）
* A_FinInst（ビジネスパートナ - 金融機関）
* A_Accounting（ビジネスパートナ - 会計）

## API への 値入力条件 の 初期値
data-platform-api-business-partner-creates-rmq-kube において、API への値入力条件の初期値は、入力ファイルレイアウトの種別毎に、次の通りとなっています。  

## データ連携基盤のAPIの選択的コール

Latona および AION の データ連携基盤 関連リソースでは、Inputs フォルダ下の sample.json の accepter に登録/更新したいデータの種別（＝APIの種別）を入力し、指定することができます。  
なお、同 accepter にAll(もしくは空白)の値を入力することで、全データ（＝全APIの種別）をまとめて登録/更新することができます。  

* sample.jsonの記載例(1)  

accepter において 下記の例のように、データの種別（＝APIの種別）を指定します。  
ここでは、"General" が指定されています。    
  
```
	"api_schema": "DPFMBusinessPartnerCreates",
	"accepter": ["General"],
```
  
* 全データを取得する際のsample.jsonの記載例(2)  

全データを取得する場合、sample.json は以下のように記載します。  

```
	"api_schema": "DPFMBusinessPartnerCreates",
	"accepter": ["All"],
```

## 指定されたデータ種別のコール

accepter における データ種別 の指定に基づいて DPFM_API_Caller 内の caller.go で API がコールされます。  
caller.go の func() 毎 の 以下の箇所が、指定された API をコールするソースコードです。  

```
func (c *DPFMAPICaller) AsyncCreates(
	accepter []string,
	input *dpfm_api_input_reader.SDC,

	log *logger.Logger,
) []error {
	wg := sync.WaitGroup{}
	mtx := sync.Mutex{}
	errs := make([]error, 0, 5)
	exconfAllExist := false

	subFuncFin := make(chan error)
	exconfFin := make(chan error)

	wg.Add(1)
	go func() {
		defer wg.Done()
		var e []error
		exconfAllExist, e = c.confirmor.Conf(input, log)
		if len(e) != 0 {
			mtx.Lock()
			errs = append(errs, e...)
			mtx.Unlock()
			exconfFin <- xerrors.Errorf("exconf error")
			return
		}
		exconfFin <- nil
	}()

	for _, fn := range accepter {
		wg.Add(1)
		switch fn {
		case "Header":
			go c.headerCreate(&wg, &mtx, subFuncFin, log, errs, input)
		case "Item":
			errs = append(errs, xerrors.Errorf("accepter Item is not implement yet"))
		default:
			wg.Done()
		}
	}
```

## Output  
本マイクロサービスでは、[golang-logging-library-for-data-platform](https://github.com/latonaio/golang-logging-library-for-data-platform) により、以下のようなデータがJSON形式で出力されます。  
以下の sample.json の例は ビジネスパートナ の 一般データ が登録/更新された結果の JSON の例です。  
以下の項目のうち、"BusinessPartner" ～ "IsMarkedForDeletion" は、/DPFM_API_Output_Formatter/type.go 内 の Type General {} による出力結果です。"cursor" ～ "time"は、golang-logging-library による 定型フォーマットの出力結果です。  

```
XXX
```
