package api

type FinancialData struct {
	ID          int    `json:"id"`
    RecDate     string `json:"recDate"`
    OpeBal      string `json:"opeBal"`
    RediscBills string `json:"rediscBills"`
    SlFacility  string `json:"slFacility"`
    SdFacility  string `json:"sdFacility"`
    Repo        string `json:"repo"`
    RevRepo     string `json:"revRepo"`
    OmoSales    string `json:"omoSales"`
    OmoRepay    string `json:"omoRepay"`
    PmSales     string `json:"pmSales"`
    PmRepay     string `json:"pmRepay"`
    Crr         string `json:"crr"`
    NetWdas     string `json:"netWdas"`
    StatAlloc   string `json:"statAlloc"`
    JvCash      string `json:"jvCash"`
    NetClr      string `json:"netClr"`
    NdicPrem    string `json:"ndicPrem"`
    OMajor      string `json:"oMajor"`
}
