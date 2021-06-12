package xmlbind

import "testing"

import (
	"encoding/xml"
	"fmt"
	"os"
)

func ExampleMarshalIndent() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	output, err := xml.MarshalIndent(v, "  ", "    ")
	if err != nil {
		fmt.Printf("error: %v\n", err)
	}

	os.Stdout.Write(output)
	// Output:
	//   <person id="13">
	//       <name>
	//           <first>John</first>
	//           <last>Doe</last>
	//       </name>
	//       <age>42</age>
	//       <Married>false</Married>
	//       <City>Hanga Roa</City>
	//       <State>Easter Island</State>
	//       <!-- Need more details. -->
	//   </person>
}

func ExampleEncoder() {
	type Address struct {
		City, State string
	}
	type Person struct {
		XMLName   xml.Name `xml:"person"`
		Id        int      `xml:"id,attr"`
		FirstName string   `xml:"name>first"`
		LastName  string   `xml:"name>last"`
		Age       int      `xml:"age"`
		Height    float32  `xml:"height,omitempty"`
		Married   bool
		Address
		Comment string `xml:",comment"`
	}

	v := &Person{Id: 13, FirstName: "John", LastName: "Doe", Age: 42}
	v.Comment = " Need more details. "
	v.Address = Address{"Hanga Roa", "Easter Island"}

	enc := xml.NewEncoder(os.Stdout)
	enc.Indent("  ", "    ")
	if err := enc.Encode(v); err != nil {
		fmt.Printf("error: %v\n", err)
	}

	// Output:
	//   <person id="13">
	//       <name>
	//           <first>John</first>
	//           <last>Doe</last>
	//       </name>
	//       <age>42</age>
	//       <Married>false</Married>
	//       <City>Hanga Roa</City>
	//       <State>Easter Island</State>
	//       <!-- Need more details. -->
	//   </person>
}

// This example demonstrates unmarshaling an XML excerpt into a value with
// some preset fields. Note that the Phone field isn't modified and that
// the XML <Company> element is ignored. Also, the Groups field is assigned
// considering the element path provided in its tag.
func ExampleUnmarshal() {
	type Email struct {
		Where string `xml:"where,attr"`
		Addr  string
	}
	type Address struct {
		City, State string
	}
	type Result struct {
		XMLName xml.Name `xml:"Person"`
		Name    string   `xml:"FullName"`
		Phone   string
		Email   []Email
		Groups  []string `xml:"Group>Value"`
		Address
	}
	v := Result{Name: "none", Phone: "none"}

	data := `
		<Person>
			<FullName>Grace R. Emlin</FullName>
			<Company>Example Inc.</Company>
			<Email where="home">
				<Addr>gre@example.com</Addr>
			</Email>
			<Email where='work'>
				<Addr>gre@work.com</Addr>
			</Email>
			<Group>
				<Value>Friends</Value>
				<Value>Squash</Value>
			</Group>
			<City>Hanga Roa</City>
			<State>Easter Island</State>
		</Person>
	`
	err := xml.Unmarshal([]byte(data), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("Name: %q\n", v.Name)
	fmt.Printf("Phone: %q\n", v.Phone)
	fmt.Printf("Email: %v\n", v.Email)
	fmt.Printf("Groups: %v\n", v.Groups)
	fmt.Printf("Address: %v\n", v.Address)
	// Output:
	// XMLName: xml.Name{Space:"", Local:"Person"}
	// Name: "Grace R. Emlin"
	// Phone: "none"
	// Email: [{home gre@example.com} {work gre@work.com}]
	// Groups: [Friends Squash]
	// Address: {Hanga Roa Easter Island}
}

func TestXmlBind(t *testing.T) {
	ExampleMarshalIndent()
	ExampleUnmarshal()
}

type SeedDoc struct {
	XMLName xml.Name `xml:"entity-engine-xml"`
	WorkloadItem []WorkloadItem `xml:"WorkloadItem"`
	WorkloadStatus []WorkloadStatus
}

type WorkloadStatus struct {
	XMLName xml.Name `xml:"WorkloadStatus"`
	StatusId string `xml:"statusId"`
}

type WorkloadItem struct{
	XMLName xml.Name `xml:"WorkloadItem"`
	WorkloadId string `xml:"workloadId,attr"`
	WorkloadItemSeqId string `xml:"workloadItemSeqId,attr"`
}

func TestSeedXmlBind(t *testing.T) {
	seed:=`<entity-engine-xml>
	<WorkloadItem workloadId="WL01" workloadItemSeqId="00001" description="WL1-001" amount="10"/>
    <WorkloadItem workloadId="WL01" workloadItemSeqId="00002" description="WL1-002" amount="20"/>
	<WorkloadStatus workloadId="WL01" statusDate="2010-01-02 00:00:00" 
		statusEndDate="2011-01-02 00:00:00" statusId="WLST_IN_DESIGN"/>
	</entity-engine-xml>`
	v:=SeedDoc{}
	err := xml.Unmarshal([]byte(seed), &v)
	if err != nil {
		fmt.Printf("error: %v", err)
		return
	}
	fmt.Printf("XMLName: %#v\n", v.XMLName)
	fmt.Printf("items: %d, %d\n", len(v.WorkloadItem), len(v.WorkloadStatus))
	for _,item := range v.WorkloadItem {
		fmt.Printf("id: %s, seq: %s\n", item.WorkloadId, item.WorkloadItemSeqId)
	}

}