<?xml version="1.0" encoding="UTF-8" standalone="yes"?>
<ProxifierProfile version="101" platform="Windows" product_id="0" product_minver="310">
  <Options>
    <Resolve>
      <AutoModeDetection enabled="false" />
      <ViaProxy enabled="true">
        <TryLocalDnsFirst enabled="false" />
      </ViaProxy>
      <ExclusionList>%ComputerName%; localhost; *.local</ExclusionList>
    </Resolve>
    <Encryption mode="basic" />
    <HttpProxiesSupport enabled="false" />
    <HandleDirectConnections enabled="false" />
    <ConnectionLoopDetection enabled="true" />
    <ProcessServices enabled="false" />
    <ProcessOtherUsers enabled="false" />
  </Options>
  <ProxyList>
    <Proxy id="100" type="SOCKS5">
      <Address>127.0.0.1</Address>
      <Port>3080</Port>
      <Options>48</Options>
    </Proxy>
  </ProxyList>
  <ChainList />
  <RuleList>
    <Rule enabled="true">
      <Name>Localhost</Name>
      <Targets>localhost; 127.0.0.1; %ComputerName%</Targets>
      <Action type="Direct" />
    </Rule>
    <Rule enabled="true">
      <Name>Direct</Name>
      <Applications>brainfuck-psiphon-pro-go.exe; </Applications>
      <Action type="Direct" />
    </Rule>
    <Rule enabled="true">
      <Name>Block</Name>
      <Applications>svchost.exe; </Applications>
      <Action type="Block" />
    </Rule>
    <Rule enabled="true">
      <Name>Default</Name>
      <Action type="Proxy">100</Action>
    </Rule>
  </RuleList>
</ProxifierProfile>
